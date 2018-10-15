import GraphiQL from "graphiql";
import React from "react";
import ReactDOM from "react-dom";

import {request as fetcher} from "./util/graphql";
import {history, Router} from "./router";

ReactDOM.render(
	<Router>
		{({response}) => {
			return <response.body params={response.params} />;
		}}
	</Router>,
	document.querySelector("#root")
);

ReactDOM.render(
	<GraphiQL fetcher={fetcher} editorTheme="ambiance" />,
	document.querySelector("#graphiql")
);
