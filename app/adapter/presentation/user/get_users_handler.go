package user

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/user"
	"github.com/kakkky/app/domain/errors"
)

type GetUsersHandler struct {
	fetchAllUserUsecase *user.ListUsersUsecase
}

func NewGetUsersHandler(fetchAllUserUsecase *user.ListUsersUsecase) *GetUsersHandler {
	return &GetUsersHandler{
		fetchAllUserUsecase: fetchAllUserUsecase,
	}
}

// @Summary     全ユーザーを取得する
// @Description 全てのユーザーのID・名前をリストで取得する
// @Tags        User
// @Produce     json
// @Success     200 {object} presenter.SuccessResponse[[]GetUsersResponse] "登録されたユーザーの情報"
// @Failure     400 {object} presenter.FailureResponse                     "不正なリクエスト"
// @Failure     500 {object} presenter.FailureResponse                     "内部サーバーエラー"
// @Router      /users [get]
func (guh *GetUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	outputs, err := guh.fetchAllUserUsecase.Run(ctx)
	if (err != nil) && errors.IsDomainErr(err) {
		presenter.RespondBadRequest(w, err.Error())
		return
	}
	if err != nil {
		presenter.RespondInternalServerError(w, err.Error())
		return
	}
	resp := make([]GetUsersResponse, 0, len(outputs))
	for _, output := range outputs {
		resp = append(resp, GetUsersResponse{
			ID:   output.ID,
			Name: output.Name,
		})
	}
	presenter.RespondOK(w, resp)

}
