// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	auth "github.com/xh-polaris/meowchat-bff/internal/handler/auth"
	collection "github.com/xh-polaris/meowchat-bff/internal/handler/collection"
	comment "github.com/xh-polaris/meowchat-bff/internal/handler/comment"
	like "github.com/xh-polaris/meowchat-bff/internal/handler/like"
	moment "github.com/xh-polaris/meowchat-bff/internal/handler/moment"
	notice "github.com/xh-polaris/meowchat-bff/internal/handler/notice"
	post "github.com/xh-polaris/meowchat-bff/internal/handler/post"
	sts "github.com/xh-polaris/meowchat-bff/internal/handler/sts"
	user "github.com/xh-polaris/meowchat-bff/internal/handler/user"
	"github.com/xh-polaris/meowchat-bff/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/account/sign_in",
				Handler: auth.SignInHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/account/send_verify_code",
				Handler: auth.SendVerifyCodeHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/account/set_password",
				Handler: auth.SetPasswordHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/notice/get_admins",
				Handler: notice.GetAdminsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/notice/get_news",
				Handler: notice.GetNewsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/notice/get_notices",
				Handler: notice.GetNoticesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notice/new_new",
				Handler: notice.NewNewsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notice/new_notice",
				Handler: notice.NewNoticeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notice/remove_new",
				Handler: notice.DeleteNewsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notice/remove_notice",
				Handler: notice.DeleteNoticeHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/like/do_like",
				Handler: like.DoLikeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/like/get_user_liked",
				Handler: like.GetUserLikedHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/like/get_count",
				Handler: like.GetLikedCountHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/post/get_post_previews",
				Handler: post.GetPostPreviewsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/post/get_post_detail",
				Handler: post.GetPostDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/post/new_post",
				Handler: post.NewPostHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/post/delete_post",
				Handler: post.DeletePostHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/post/search_post",
				Handler: post.SearchPostHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/collection/get_cat_previews",
				Handler: collection.GetCatPreviewsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/collection/get_cat_detail",
				Handler: collection.GetCatDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/collection/new_cat",
				Handler: collection.NewCatHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/collection/delete_cat",
				Handler: collection.DeleteCatHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/post/search_cat",
				Handler: collection.SearchCatHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/moment/get_moment_previews",
				Handler: moment.GetMomentPreviewsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/moment/get_moment_detail",
				Handler: moment.GetMomentDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/moment/new_moment",
				Handler: moment.NewMomentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/moment/delete_moment",
				Handler: moment.DeleteMomentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/post/search_moment",
				Handler: moment.SearchMomentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/comment/get_comments",
				Handler: comment.GetCommentsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/comment/new_comment",
				Handler: comment.NewCommentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/get_user_info",
				Handler: user.GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/update_user_info",
				Handler: user.UpdateUserInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/sts/apply_token",
				Handler: sts.ApplyTokenHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
