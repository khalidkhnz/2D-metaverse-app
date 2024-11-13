import { ENDPOINTS } from "@/lib/Endpoints";
import { getMethod, postMethod } from "../ApiInterceptor";
import { BODY } from "@/lib/Forms";
import { ISpaceApiResponse } from "@/types/space";

export class SpaceService {
  public static async handleCreateSpace(spaceName: string) {
    try {
      const response = await postMethod(
        ENDPOINTS.SPACE.CREATE_SPACE,
        BODY.SPACE.CREATE_SPACE(spaceName),
      );
      return response;
    } catch (error) {
      console.log("error-SpaceService-handleCreateSpace");
    }
  }

  public static async handleGetAllSpces() {
    try {
      const response = await getMethod(ENDPOINTS.SPACE.GET_ALL_SPACES);
      return response;
    } catch (error) {
      console.log("error-SpaceService-handleGetAllSpces");
    }
  }

  public static async handleGetSpaceById(_id: string) {
    try {
      const response = await getMethod(ENDPOINTS.SPACE.GET_SPACES_BY_ID(_id));
      return response;
    } catch (error) {
      console.log("error-SpaceService-handleGetSpaceById");
    }
  }

  public static async handleGetAllMySpace() {
    try {
      const response = await getMethod<ISpaceApiResponse>(
        ENDPOINTS.SPACE.GET_ALL_MY_SPACES,
      );
      return response;
    } catch (error) {
      console.log("error-SpaceService-handleGetSpaceById");
    }
  }

  public static async handleSearchSpace(searchTerm: string) {
    try {
      const response = await getMethod(
        ENDPOINTS.SPACE.SEARCH_SPACE(searchTerm),
      );
      return response;
    } catch (error) {
      console.log("error-SpaceService-handleSearchSpace");
    }
  }
}
