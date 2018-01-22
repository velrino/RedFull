/**
* @api {get} /api/users List
* @apiName List Users
* @apiGroup Users
*
* @apiSuccess {String} name Name of the User.
* @apiSuccess {String} email E-mail of the User.
* @apiSuccess {String} gravatar Gravatar of the User.
* @apiSampleRequest /api/users
* @apiHeader {String} Authorization <code>Bearer {TOKEN}</code> Bearer token for authentication.
* @apiSuccessExample Success-Response:
HTTP/1.1 200 OK
[
	{
		"id": 1,
		"email": "colin@redventures.com",
		"gravatar": "http://www.Gravatar.com/avatar/a51972ea936bc3b841350caef34ea47e?s=64&d=monsterID",
		"name": "Colin"
	},
	{
		"id": 2,
		"email": "kyle@redventures.com",
		"gravatar": "http://www.Gravatar.com/avatar/432f3e353c689fc37af86ae861d934f9?s=64&d=monsterID",
		"name": "Kyle"
	},
]
* @apiError UserNotFound The Users was not found.
*
* @apiErrorExample Error-Response:
HTTP/1.1 404 Not Found
{
	"message": "Users Not Found"
}
*/

/**
* @api {get} /api/users/:id Get
* @apiName Get Users
* @apiGroup Users
*
* @apiSuccess {String} name Name of the User.
* @apiSuccess {String} email E-mail of the User.
* @apiSuccess {String} gravatar Gravatar of the User.
* @apiSampleRequest /api/users/1
* @apiHeader {String} Authorization <code>Bearer {TOKEN}</code> Bearer token for authentication.
* @apiSuccessExample Success-Response:
HTTP/1.1 200 OK
{
	"id": 1,
	"email": "colin@redventures.com",
	"gravatar": "http://www.Gravatar.com/avatar/a51972ea936bc3b841350caef34ea47e?s=64&d=monsterID",
	"name": "Colin"
}
* @apiError UserNotFound The Users was not found.
*
* @apiErrorExample Error-Response:
HTTP/1.1 404 Not Found
{
	"message": "User Not Found"
}
*/


/**
* @api {get} /api/widgets/ List
* @apiName List Widget
* @apiGroup Widgets
*
* @apiSuccess {String} id Id of the Widget.
* @apiSuccess {String} name Name of the Widget.
* @apiSuccess {String} inventory Inventory of the Widget.
* @apiSuccess {String} color Color of the Widget.
* @apiSuccess {String} price Price of the Widget.
* @apiSuccess {String} melts Melts of the Widget.
* @apiSampleRequest /api/widgets
* @apiHeader {String} Authorization <code>Bearer {TOKEN}</code> Bearer token for authentication.
* @apiSuccessExample Success-Response:
* HTTP/1.1 200 OK
[
    {
        "id": 1,
        "color": "blue",
        "inventory": 51,
        "name ": "Losenoid",
        "price": "9.99"
    },
    {
        "id": 2,
        "color": "red",
        "inventory": 7,
        "name ": "Rowlow",
        "price": "4.00"
    }
]
* @apiError WidgetsNotFound The Widgets was not found.
*
* @apiErrorExample Error-Response:
*     HTTP/1.1 404 Not Found
*     {
*       "message": "Widgets Not Found"
*     }
*/

/**
* @api {get} /api/widgets/:id Get
* @apiName Get Widget
* @apiGroup Widgets
* @apiSuccess {String} id Id of the Widget.
* @apiSuccess {String} name Name of the Widget.
* @apiSuccess {String} inventory Inventory of the Widget.
* @apiSuccess {String} color Color of the Widget.
* @apiSuccess {String} price Price of the Widget.
* @apiSuccess {String} melts Melts of the Widget.
* @apiSampleRequest /api/widgets/1
* @apiHeader {String} Authorization <code>Bearer {TOKEN}</code> Bearer token for authentication.
* @apiSuccessExample Success-Response:
HTTP/1.1 200 OK
{
	"id": 1,
	"color": "blue",
	"inventory": 51,
	"name ": "Losenoid",
	"price": "9.99"
}
* @apiError WidgetNotFound The Widget was not found.
*
* @apiErrorExample Error-Response:
HTTP/1.1 404 Not Found
{
	"message": "Widget Not Found"
}
*/

/**
* @api {post} /api/widgets Create
* @apiName Create Widget
* @apiGroup Widgets
* @apiParam {String} Name Name of the Widget.
* @apiParam {Number} Inventory Inventory of the Widget.
* @apiParam {String} Color Color of the Widget.
* @apiParam {Number} Price Price of the Widget.
* @apiParam {Boolean} Melts Melts of the Widget.
* @apiSampleRequest /api/widgets
* @apiHeader {String} Authorization <code>Bearer {TOKEN}</code> Bearer token for authentication.
* @apiSuccessExample Success-Response:
HTTP/1.1 200 OK
{
	"message": "Widget created successfully!"
}
* @apiError WidgetBadRequest The Widget bad request.
* @apiErrorExample Error-Response:
HTTP/1.1 400 Bad Request
{
	"message": "Incorrect param"
}
*/

/**
* @api {put} /api/widgets/:id Update
* @apiName Update Widget
* @apiGroup Widgets
* @apiParam {String} name Name of the Widget.
* @apiParam {Number} inventory Inventory of the Widget.
* @apiParam {String} color Color of the Widget.
* @apiParam {Number} price Price of the Widget.
* @apiParam {Boolean} melts Melts of the Widget.
* @apiSampleRequest /api/widgets/1
* @apiHeader {String} Authorization <code>Bearer {TOKEN}</code> Bearer token for authentication.
* @apiSuccessExample Success-Response:
HTTP/1.1 200 OK
{
	"message": "Widget updated successfully!"
}
* @apiError WidgetNotFound The Widget was not found.
* @apiErrorExample Error-Response:
HTTP/1.1 404 Not Found
{
	"message": "Widget updated successfully!"
}
* @apiError WidgetBadRequest The Widget bad request.
* @apiErrorExample Error-Response:
HTTP/1.1 400 Bad Request
{
	"message": "Incorrect param"
}
*/

/**
* @api {post} /login Login
* @apiName Login
* @apiGroup Auth
* @apiHeader {String} Content-Type <code>application/json</code> JSON representation
* @apiParam {String} username E-mail of user
* @apiParam {String} password Password default <code>12345</code>
* @apiSampleRequest /login
*/