package rest_errors

import (
	"net/http"
)

var (
	// Common
	ErrCopyStruct    = NewError("failed to copy the structs", http.StatusInternalServerError)
	ErrMissingParams = NewError("params not found in the request", http.StatusBadRequest)

	// Controller
	ErrParsingRequestBody   = NewError("failed to parse request body", http.StatusBadRequest)
	ErrParsingRequestParams = NewError("failed to parse request params", http.StatusBadRequest)
	NoLoggedInUserFound     = NewError("no logged-in user found", http.StatusUnauthorized)
	AccessForbidden         = NewError("access forbidden", http.StatusForbidden)

	// Service
	ErrLogin                      = NewError("invalid email or password", http.StatusUnauthorized)
	ErrDeletedUserAccess          = NewError("the account has been removed by the owner", http.StatusForbidden)
	ErrLogOut                     = NewError("failed to logout", http.StatusInternalServerError)
	ErrSamePassword               = NewError("password can't be same as old one", http.StatusBadRequest)
	ErrResendingForgotPasswordOTP = NewError("failed to resend the OTP", http.StatusInternalServerError)
	ErrCreatingForgotPasswordOTP  = NewError("failed to create the OTP", http.StatusInternalServerError)
	ErrUsingPrivateEmail          = NewError("you cannot use a hidden email to sign in.", http.StatusUnauthorized)
	ErrFetchingEmail              = NewError("failed to fetch email from provider", http.StatusUnauthorized)

	InvalidSigningMethod  = NewError("invalid signing method while parsing jwt", http.StatusUnauthorized)
	InvalidPasswordFormat = NewError("minimum 8 characters with at least 1 uppercase letter(A-Z), 1 lowercase letter(a-z), 1 number(0-9) and 1 special character(.!@#~$%^&*()+|_<>)", http.StatusBadRequest)

	// Report
	ReportedContentNotFound    = NewError(NotFound("reported content"), http.StatusNotFound)
	ErrUpdatingReportedContent = NewError(Update("reported content"), http.StatusInternalServerError)

	ErrReportAlreadyCreated = NewError("report is already created", http.StatusUnprocessableEntity)

	// Report
	ErrCreatingBlock       = NewError("failed to block the user", http.StatusInternalServerError)
	ErrBlockAlreadyCreated = NewError("you have already blocked the user", http.StatusUnprocessableEntity)
	ErrBlockingOwnAccount  = NewError("you cannot block yourself", http.StatusBadRequest)
	ErrUnblockingUser      = NewError("failed to unblock the user", http.StatusInternalServerError)
	BlockedUserNotFound    = NewError(NotFound("blocked user"), http.StatusNotFound)

	// validation errors
	ErrEmailUpdateNotAllowed     = NewError("email update not allowed", http.StatusBadRequest)
	ErrUserNameUpdateNotAllowed  = NewError("username update not allowed", http.StatusBadRequest)
	ErrPasswordUpdateNotAllowed  = NewError("password update not allowed", http.StatusBadRequest)
	ErrInvalidLoginProvider      = NewError("invalid login provider", http.StatusBadRequest)
	ErrEmailAlreadyRegistered    = NewError("email is already taken", http.StatusBadRequest)
	ErrUserNameAlreadyRegistered = NewError("username is already taken", http.StatusBadRequest)
)
