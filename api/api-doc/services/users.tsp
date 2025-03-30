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
@service(#{ title: "Users API" })
@server(ProductionServer, "Production")
@server(StagingServer, "Staging")
@server(DevelopmentServer, "Development")
@server(LocalhostServer, "Localhost")
namespace UsersApi {
  model User {
    id: string;
    name: string;
    email: string;
    createdAt: utcDateTime;
    updatedAt: utcDateTime;
  }

  model GetUserByIdResponse {
    user: User;
  }

  model GetUsersResponse {
    users: User[];
  }

  model UpdateUserRequest {
    name: string;
    email: string;
    password: string;
  }
    
  @route("/users")
  @tag("Users")
  interface Users {
    @doc("Delete an user by the id")
    @useAuth(BearerAuth)
    @route("/{id}")
    @delete
    deleteUserById(
      @path id: string
    ): ResponseDeleted<null> 
      | CommonClientResponseErrors
      | CommonServerResponseErrors 
      | ResponseUnauthorizedError 
      | ResponseForbiddenError 
      | ResponseNotFoundError 
      | ResponseValidationError;

    @doc("Get an user by the id")
    @useAuth(BearerAuth)
    @route("/{id}")
    @get
    getUserById(
      @path id: string
    ): ResponseRead<GetUserByIdResponse> 
      | CommonClientResponseErrors
      | CommonServerResponseErrors 
      | ResponseUnauthorizedError 
      | ResponseForbiddenError 
      | ResponseNotFoundError 
      | ResponseValidationError;
  
    @doc("List all users")
    @useAuth(BearerAuth)
    @get
    getUsers(
      @query query?: PaginationQuery<{ name: string }>
    ): ResponseRead<GetUsersResponse>
      | CommonClientResponseErrors
      | CommonServerResponseErrors
      | ResponseUnauthorizedError
      | ResponseForbiddenError
      | ResponseNotFoundError;

    @doc("Update an user")
    @useAuth(BearerAuth)
    @route("/{id}")
    @patch
    updateUser(
      @header("Content-Type") contentType: "application/json";
      @path id: string,
      @body updateUserRequest: UpdateUserRequest
    ): ResponseUpdated<null>
      | CommonClientResponseErrors
      | CommonServerResponseErrors
      | ResponseUnauthorizedError
      | ResponseForbiddenError
      | ResponseNotFoundError
      | ResponseConflictError
      | ResponseValidationError;
  }
}
