export namespace entity {
	
	export class Option {
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    // Go type: gorm
	    deletedAt: any;
	    version: number;
	    id: number;
	    key: string;
	    value: string;
	    valueEncrypted: string;
	    encrypted: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Option(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.deletedAt = this.convertValues(source["deletedAt"], null);
	        this.version = source["version"];
	        this.id = source["id"];
	        this.key = source["key"];
	        this.value = source["value"];
	        this.valueEncrypted = source["valueEncrypted"];
	        this.encrypted = source["encrypted"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

