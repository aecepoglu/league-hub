import React from "react";
import ReactDOM from "react-dom";

import {history, Router} from "./router";

ReactDOM.render(
	<Router>
		{({response}) => {
			return <response.body params={response.params} />;
		}}
	</Router>,
	document.querySelector("#root")
);
