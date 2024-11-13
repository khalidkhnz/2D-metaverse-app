export interface ISpace {
  _id: string;
  name: string;
  description: string;
  logo: string;
  coverArt: string;
  creatorId: string;
  memberIds: string[];
  adminIds: string[];
  moderatorIds: string[];
  elderIds: string[];
  channelIds: string[];
  createdAt: string;
  updatedAt: string;
}

export interface ISpaceApiResponse {
  success: boolean;
  data: ISpace[];
}
