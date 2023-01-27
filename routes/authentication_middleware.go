package routes

import (
	"context"
	"net/http"
	"social-media-api/service"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc/metadata"
)

func authenticationMiddleware(svc *service.Services) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username := mux.Vars(r)["username"]
			userID, err := svc.TwitterService.AuthenticateUser(r.Context(), username)
			if userID == "" || err != nil {
				http.Error(w, "[AuthenticationMiddleware] Invalid username", http.StatusBadRequest)
				return
			}
			//Inject user id into context so user doesn't have to bother with finding user ids
			meta := make(map[string]string)
			meta["userID"] = userID
			ctx := metadata.NewOutgoingContext(r.Context(), metadata.New(meta))
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
