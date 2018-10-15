import Browser from "@hickory/browser";
import {curi} from "@curi/router";
import {curiProvider} from "@curi/react-dom";

import Login from "./login.jsx";
import Home from "./home.jsx";
import NotFound from "./not-found.jsx";
import Welcome from "./welcome.jsx";

const ROUTES = [{
	name: "welcome",
	path: "",
	response: () => ({body: Welcome})
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
