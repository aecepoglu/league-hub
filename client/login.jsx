import React from "react";
import {send as gql} from "./util/graphql";
import {router} from "./router";
import auth from "./util/auth";

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
		e.preventDefault();
		gql("query($U: String!, $P: String!) { login(email: $U, password: $P) { token } }", {
			U: this.state.email.current.value,
			P: this.state.password.current.value
		}).then(x => {
			auth.save(x.login.token);
			router.navigate({name: "home"});
		}).catch(e => {
			console.log("bad", e);
		});
	}
}

export default Login;