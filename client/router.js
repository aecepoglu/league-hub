import Browser from "@hickory/browser";
import {curi} from "@curi/router";
import {curiProvider} from "@curi/react-dom";

import Login from "./components/login";
import Home from "./components/home";
import NotFound from "./components/not-found";
import Welcome from "./components/welcome";
import AdminPanel from "./components/admin-panel";

const ROUTES = [{
	name: "welcome",
	path: "",
	response: () => ({body: Welcome})
}, {
	name: "home",
	path: "home",
	response: () => ({body: Home})
}, {
	name: "admin",
	path: "admin",
	response: () => ({body: AdminPanel})
}, {
	name: "not found",
	path: "(.*)",
	response: () => ({body: NotFound})
}];

const history = Browser();

const router = curi(history, ROUTES);
const Router = curiProvider(router);

export {history, router, Router};
