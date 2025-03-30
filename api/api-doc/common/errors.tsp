import "@typespec/http";
import "./common.tsp";

using TypeSpec.Http;
using Common;

namespace Errors {
  @error
  model ResponseError<StatusCode, ErrorDetails> {
    @body body: ApiResponse<{}, ErrorDetails>;
    @statusCode status: StatusCode;
  }

  // 400
  model ResponseBadRequestErrorDetails {}

  @error
  model ResponseBadRequestError {
    ...ResponseError<400, ResponseBadRequestErrorDetails>;
  }

  // 401
  model ResponseUnauthorizedErrorDetails {}

  @error
  model ResponseUnauthorizedError {
    ...ResponseError<401, ResponseUnauthorizedErrorDetails>;
  }

  // 403
  model ResponseForbiddenErrorDetails {}

  @error
  model ResponseForbiddenError {
    ...ResponseError<403, ResponseForbiddenErrorDetails>;
  }

  // 404
  model ResponseNotFoundErrorDetails {
    @doc("Resource not found")
    resource: string;
  }

  @error
  model ResponseNotFoundError {
    ...ResponseError<404, ResponseNotFoundErrorDetails>;
  }

  // 406
  model ResponseNotAcceptableErrorDetails {
    @doc("Content types that the server produces")
    contentTypes: string[];
  }

  @error
  model ResponseNotAcceptableError {
    ...ResponseError<406, ResponseNotAcceptableErrorDetails>;
  }

  // 409
  model ResponseConflictErrorDetails {
    @doc("Resource in conflict")
    resource: string;
  }

  @error
  model ResponseConflictError {
    ...ResponseError<409, ResponseConflictErrorDetails>;
  }

  // 414
  model ResponseUriTooLongErrorDetails {}

  @error
  model ResponseUriTooLongError {
    ...ResponseError<414, ResponseUriTooLongErrorDetails>;
  }

  // 415
  model ResponseUnsupportedMediaTypeErrorDetails {
    @doc("Content types that the server accepts")
    contentTypes: string[];
  }

  @error
  model ResponseUnsupportedMediaTypeError {
    ...ResponseError<415, ResponseUnsupportedMediaTypeErrorDetails>;
  }

  // 422
  model ResponseValidationErrorDetails {
    field: string;
    message: string;
  }
  
  @error
  model ResponseValidationError {
    ...ResponseError<422, ResponseValidationErrorDetails>;
  }

  // 429
  model ResponseRateLimitErrorDetails {
    waitingPeriodInSeconds: uint32;
  }
  
  @error
  model ResponseRateLimitError {
    ...ResponseError<429, ResponseRateLimitErrorDetails>;
  }

  alias CommonClientResponseErrors =
    ResponseBadRequestError |
    ResponseNotAcceptableError |
    ResponseUriTooLongError |
    ResponseUnsupportedMediaTypeError |
    ResponseRateLimitError;

  // 500
  model ResponseInternalServerErrorDetails {}
  
  @error
  model ResponseInternalServerError {
    ...ResponseError<500, ResponseInternalServerErrorDetails>;
  }

  // 502
  model ResponseBadGatewayErrorDetails {
    resource: string
  }
  
  @error
  model ResponseBadGatewayError {
    ...ResponseError<502, ResponseBadGatewayErrorDetails>;
  }

  alias CommonServerResponseErrors = ResponseInternalServerError;
}
