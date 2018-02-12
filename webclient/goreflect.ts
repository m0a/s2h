export type GoReflectKind = "ptr" | "struct" | "string" | string;

export default interface GoReflect {
  kind: GoReflectKind;
  type?: string;
  value?: string;
  order: number;
  fields: { [key: string]: GoReflect };
};
