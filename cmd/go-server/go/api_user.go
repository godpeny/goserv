/*
 * goserv
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	ent_client "github.com/godpeny/goserv/internal/clients/ent"
	"github.com/godpeny/goserv/pkg/types"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	ctx := r.Context()
	log.Ctx(ctx).Info().Msgf("CreateUser")

	req := &V1User{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Error().Interface("CreateUser", req).Msgf("error from CreateUser decodeing process %s", err)
		render.Render(w, r, types.ErrNotFound)
		return
	}

	// db
	client, err := ent_client.GetClient()
	if err != nil {
		log.Error().Msgf("error from CreateUser get ent client %s", err)
	}

	_, err = client.User.Create().SetAge(int(req.Age)).SetName(req.Name).Save(ctx)
	if err != nil {
		log.Error().Interface("CreateUser", req).Msgf("error from CreateUser db %s", err)
		render.Render(w, r, types.ErrRender(err))
		return
	}

	log.Info().Msgf("%v", fmt.Sprintf("%v", req))

	log.Ctx(ctx).Info().Msgf("response from CreateUser")
	render.JSON(w, r, req)
}
