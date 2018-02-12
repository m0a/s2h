export default interface GoReflect {
  kind: string;
  type?: string;
  value?: string;
  order: number;
  fields: { [key: string]: GoReflect };
};
