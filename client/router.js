import Browser from "@hickory/browser";
import {curi} from "@curi/router";
import {curiProvider} from "@curi/react-dom";

import Login from "./login.jsx";
import Home from "./home.jsx";
import NotFound from "./not-found.jsx";

const ROUTES = [{
	name: "login",
	path: "",
	response: () => ({body: Login})
}, {
	name: "home",
	path: "home",
	response: () => ({body: Home})
}, {
	name: "not found",
	path: "(.*)",
	response: () => ({body: NotFound})
}];

const history = Browser();

const router = curi(history, ROUTES);
const Router = curiProvider(router);

export {history, router, Router};
