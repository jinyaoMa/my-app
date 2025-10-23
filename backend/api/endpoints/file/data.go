package file

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/api/apir"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type DataInput struct {
	id datatype.Id
	Id string `path:"id" required:"true" doc:"File Id (Base36)"`
}

func (i *DataInput) Resolve(ctx huma.Context) (errs []error) {
	id, err := datatype.ParseIdFromB36(i.Id)
	if err != nil {
		errs = append(errs, err)
		return
	}

	i.id = id
	return
}

func (a *File) RegisterData(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:       http.MethodGet,
		Path:         "/file/data/{id}",
		Summary:      "Get File Data",
		Description:  "Get file data by id.",
		OperationID:  "file.get.data",
		Tags:         []string{"file"},
		Errors:       []int{http.StatusInternalServerError},
		MaxBodyBytes: 1,
	}

	op.Security = append(op.Security, map[string][]string{
		apir.AuthScheme: {op.OperationID},
	})

	handler := func(ctx context.Context, input *DataInput) (output *huma.StreamResponse, err error) {
		service, cancel := service.NewFileService(ctx, a.CrudOptions, a.Db, a.Store)
		defer cancel()

		file, mime, err := service.OpenFileById(input.id)
		if err != nil {
			return nil, err
		}

		return &huma.StreamResponse{
			Body: func(ctx huma.Context) {
				defer file.Close()

				// Write header info before streaming the body.
				if mime == "" {
					mime = "application/octet-stream"
				}
				ctx.SetHeader("Content-Type", mime)
				writer := ctx.BodyWriter()

				// Update the write deadline to give us extra time.
				if d, ok := writer.(interface{ SetWriteDeadline(time.Time) error }); ok {
					d.SetWriteDeadline(time.Now().Add(12 * time.Hour))
				}

				buffer := make([]byte, 64)
				if f, ok := writer.(http.Flusher); ok {
					// Write, then flush and wait.
					for {
						select {
						case <-ctx.Context().Done():
							return
						default:
							n, errRead := file.Read(buffer)
							if errRead == io.EOF {
								return
							}
							if errRead != nil {
								return
							}

							m, errWrite := writer.Write(buffer[:n])
							if errWrite != nil {
								return
							}
							if n != m {
								return
							}
							f.Flush()
						}
					}
				}
			},
		}, nil
	}

	huma.Register(api, op, handler)
	return
}
