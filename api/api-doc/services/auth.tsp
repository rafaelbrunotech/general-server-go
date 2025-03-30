import "@typespec/http";
import "@typespec/versioning";
import "../common/common.tsp";
import "../common/constants.tsp";
import "../common/errors.tsp";

using TypeSpec.Http;
using TypeSpec.Versioning;
using Common;
using Constants;
using Errors;

@versioned(Versions)
@service(#{ title: "Auth API" })
@server(ProductionServer, "Production")
@server(StagingServer, "Staging")
@server(DevelopmentServer, "Development")
@server(LocalhostServer, "Localhost")
namespace AuthApi {
  model AuthUser {
    id: string;
    name: string;
    accessToken: string;
    refreshToken: string;
  }
  
  model SignInRequest {
    email: string;
    password: string;
  }

  model SignInResponse {
    authUser: AuthUser;
  }
  
  model SignUpRequest {
    name: string;
    email: string;
    password: string;
  }

  model SignUpResponse {
    authUser: AuthUser;
  }

  @tag("Auth")
  interface Auth {
    @doc("Sign the user in")
    @route("/sign-in")
    @post
    signIn(
      @header("Content-Type") contentType: "application/json";
      @body signInRequest: SignInRequest
    ): ResponseRead<SignInResponse>
      | CommonClientResponseErrors
      | CommonServerResponseErrors
      | ResponseUnauthorizedError
      | ResponseNotFoundError
      | ResponseValidationError;

    @doc("Sign the user up")
    @route("/sign-up")
    @post
    signUp(
      @header("Content-Type") contentType: "application/json";
      @body signUpRequest: SignUpRequest
    ): ResponseCreated<SignUpResponse, 200>
      | CommonClientResponseErrors
      | CommonServerResponseErrors
      | ResponseConflictError
      | ResponseValidationError;
  }
}
