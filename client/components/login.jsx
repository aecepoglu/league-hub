import React from "react";
import AsyncButton from 'react-async-button';
import {send as gql} from "../util/graphql";
import {router} from "../router";
import auth from "../util/auth";

class Login extends React.Component {
	constructor(props) {
		super(props);
		this.handleSubmit = this.handleSubmit.bind(this);
		this.noop = e => e.preventDefault();
		this.state = {
			email: React.createRef(),
			password: React.createRef(),
			errors: []
		};
	}

	render() {
		return (
			<form className="box is-half" onSubmit={this.noop}>
				<div className="title">Log-In</div>

				<div className="field">
					<label className="label">E-Mail</label>
					<div className="control has-icons-left">
						<input className="input" type="text" placeholder="email" ref={this.state.email} />
						<span className="icon is-small is-left">
							<i className="fas fa-envelope"></i>
						</span>
					</div>
				</div>

				<div className="field">
					<label className="field">Password</label>
					<div className="control has-icons-left">
						<input className="input" type="password" ref={this.state.password} />
						<span className="icon is-small is-left">
							<i className="fas fa-key"></i>
						</span>
					</div>

					{this.state.errors.map((e, i) => (
						<p key={i} className="help is-danger is-size-6">{e.message}</p>
					))}
				</div>

				<div className="field">

					<div className="control">
						<AsyncButton className="button is-link"
							loadingClass="is-loading"
							text="OK"
							onClick={this.handleSubmit}>
						</AsyncButton>
					</div>
				</div>
			</form>
		)
	}

	handleSubmit(e) {
		e.preventDefault();
		return gql("query($U: String!, $P: String!) { login(email: $U, password: $P) { token } }", {
			U: this.state.email.current.value,
			P: this.state.password.current.value
		}).then(x => {
			auth.save(x.login.token);
			router.navigate({name: "home"});
		}).catch(e => {
			this.setState({errors: e});
		}).then(() => {
			return new Promise(resolve => {
				setTimeout(resolve, 250);
			});
		});
	}
}

export default Login;
