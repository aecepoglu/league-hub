import React from "react";
import GraphiQL from "graphiql";
import {request as fetcher} from "../../util/graphql";
import style from "./style.css";

class AdminPanel extends React.Component {
	render() {
		return (
			<div>
				<p className="title">Admin Panel</p>

				<p className="subtitle">GraphiQL</p>
				<div className={style.myContainer}>
					<GraphiQL fetcher={fetcher} editorTheme="ambiance" />
				</div>
			</div>
		)
	}
}

export default AdminPanel
