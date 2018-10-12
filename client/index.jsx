const name = "ahmet emre";

const React = require("react");
const ReactDOM = require("react-dom");
const GraphiQL = require("graphiql");

const Login = require("./login.jsx");
import {Fetch as fetcher} from "./util/graphql";

ReactDOM.render(
	<Login />,
	document.querySelector("#root")
);

ReactDOM.render(
	<GraphiQL fetcher={fetcher} editorTheme="ambiance" />,
	document.querySelector("#graphiql")
);
