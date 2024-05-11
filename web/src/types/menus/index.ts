export interface menuType {
  pageID: string;
  pageName: string;
  pagePath: string;
  pageIcon: string;
  pageComponent: string;
  parentPage: string;
  children: menuType[];
  pageOrder: number;
  canEdit: 1 | 0;
  isOutSite: boolean;
  outSiteLink: string;
  createdAt: string;
  updateTime: string;
}
