import { Ref } from "vue";

export type Type =
  | "primary"
  | "success"
  | "warning"
  | "danger"
  | "info"
  | "default";
export type TypeRef = Ref<Type>;

export type Size = "small" | "large" | "default";
export type SizeRef = Ref<Size>;

export type Direction = "vertical" | "horizontal";
export type DirectionRef = Ref<Direction>;

export type LabelPosition = "left" | "top";
export type LabelPositionRef = Ref<LabelPosition>;

export type ResponseState = "success" | "warning" | "error";
