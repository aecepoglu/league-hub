const name = "ahmet emre";

function gqlRequest(query, variables) {
	return fetch(window.location.origin + "/graphql", {
		method: "POST",
		headers: {
			"content-type": "application/json"
		},
		body: JSON.stringify({
			query: query,
			variables: variables
		})
	}).then(resp => resp.json())
		.then(x => {
			if (x.errors) {
				return Promise.reject(x.errors);
			}
			return x.data;
		});
}

function fetcher(params) {
	return gqlRequest(params.query, params.variables);
}

class Login extends React.Component {
	constructor(props) {
		super(props);
		this.handleSubmit = this.handleSubmit.bind(this);
		this.state = {
			email: React.createRef(),
			password: React.createRef()
		};
	}

	render() {
		return (
			<form className="box is-half" onSubmit={this.handleSubmit}>
				<div className="title">Log-In</div>

				<div className="field">
					<label className="label">E-Mail</label>
					<div className="control has-icons-left">
						<input className="input" type="text" placeholder="email" ref={this.state.email}/>
						<span className="icon is-small is-left">
							<i className="fas fa-envelope"></i>
						</span>
					</div>
				</div>

				<div className="field">
					<label className="field">Password</label>
					<div className="control has-icons-left">
						<input className="input" type="password" ref={this.state.password}/>
						<span className="icon is-small is-left">
							<i className="fas fa-key"></i>
						</span>
					</div>
				</div>

				<div className="field">
					<div className="control">
						<button className="button is-link">OK</button>
					</div>
				</div>
			</form>
		)
	}

	handleSubmit(e) {
		gqlRequest("query($U: String!, $P: String!) { login(email: $U, password: $P) { token } }", {
			U: this.state.email.current.value,
			P: this.state.password.current.value
		}).then(x => {
			console.log("good", x);
		}).catch(e => {
			console.log("bad", e);
		});
		e.preventDefault();
	}
}

ReactDOM.render(
	<Login />,
	root
);

ReactDOM.render(
	<GraphiQL fetcher={gqlRequest} editorTheme="ambiance" />,
	document.querySelector("#graphiql")
);
