export type TreePath = string | string[];

export type ItemHighlight = {
  name: string;
  hl_group: string;
  col: number;
  width: number;
};

export type ItemStatus = {
  size?: number;
  time?: number;
};

export type Item<
  ActionData extends unknown = unknown,
> = {
  word: string;
  display?: string;
  action?: ActionData;
  data?: unknown;
  highlights?: ItemHighlight[];
  status?: ItemStatus;
  kind?: string;
  level?: number;
  treePath?: TreePath;
  isExpanded?: boolean;
  isTree?: boolean;
};

export type DduActionData = unknown;

export type DduItem =
  & Item<DduActionData>
  & {
    matcherKey: string;
    __sourceIndex: number;
    __sourceName: string;
    __level: number;
    __expanded: boolean;
    __columnTexts: Record<number, string>;
    __groupedPath: string;
  };
