import "@typespec/http";
import "./constants.tsp";

using TypeSpec.Http;
using Constants;

namespace Common {
  model ApiResponseMetadata {
    environment: Environment;
    page: uint32 | null;
    perPage: uint32 | null;
    total: uint32 | null;
  }

  model ApiResponseError<E> {
    details: E;
    message: string;
  }

  model ApiResponse<T, E> {
    data: T;
    error: ApiResponseError<E> | null;
    metadata: ApiResponseMetadata;
    status: uint16;
    timestamp: utcDateTime;
  }
  
  model ResponseCreated<T, Status = 204> {
    @body body: ApiResponse<T, null>;
    @statusCode status: Status;
  }

  model ResponseRead<T> {
    @body body: ApiResponse<T, null>;
    @statusCode status: 200;
  }
  
  model ResponseUpdated<T> {
    @body body: ApiResponse<T, null>;
    @statusCode status: 204;
  }

  model ResponseDeleted<T> {
    @body body: ApiResponse<T, null>;
    @statusCode status: 200;
  }

  model PaginationQuery<Filter = {}> {
    @pageIndex page: uint32;
    @pageSize perPage: uint32;
    
    ...Filter;
  }
}