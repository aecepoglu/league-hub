import React from "react";
import AsyncButton from "react-async-button";
import {send as gql} from "../util/graphql";
import ValidationError from "../errs/validation";
import ErrMsgs from "./err-msgs.jsx";

const REQUIRED = {
	name: "required",
	fun: function(x) {
		return x === undefined ||
			x == "";
	}
};

function validateSubmissionForm(state, rules) {
	let errors = {};

	Object.keys(rules).forEach(key => {
		let obj = {};

		rules[key].forEach(r => {
			obj[r.name] = r.fun(state[key])
		});

		if (Object.values(obj).every(x => x)) {
			errors[key] = obj;
			errors.$ = true;
		}
	});


	return errors;
}

class LeagueCreator extends React.Component {
	constructor(props) {
		super(props);

		this.sports = [
			"Squash",
			"Tennis",
		];

		this.state = {
			name: "my league",
			type: "",
			isPrivate: false,
			isOpen: true,
			errors: {},
		};

		this.RULES = {
			name: [REQUIRED],
			type: [REQUIRED],
		};

		this.submit = this.submit.bind(this);
		this.handleInputChange = this.handleInputChange.bind(this);
		this.handleRadioChange = ev => this.handleInputChange({
			target: {
				name: ev.target.name,
				value: ev.target.value === "yes"
			}
		});
		this.noop = e => e.preventDefault();
	}

	handleInputChange(ev) {
		this.setState({
			[ev.target.name]: ev.target.value
		});
	}

	submit() {
		return Promise.resolve().then(() => {
			let errors = validateSubmissionForm(this.state, this.RULES);
			if (errors.$) {
				throw new ValidationError(errors);
			}

			console.log(this.state);
		}).catch(err => {
			if (err instanceof ValidationError) {
				this.setState({errors: err.errors});
			}
		});
	}

	render() {
		return (
			<form onSubmit={this.noop}>
				<div className="field">
					<div className="control">
						<input className="input" type="text" name="name"
							value={this.state.name} onChange={this.handleInputChange}
							placeholder="Name"
						/>

						<ErrMsgs className="help is-danger" for={this.state.errors.name}
							hiddenClass="is-hidden"
						>
							<span for="required">This field is required</span>
							<span for="unique">Name is already in use</span>
						</ErrMsgs>
					</div>
				</div>
				<div className="field">
					<div className="control">
						<div className="select">
							<select value={this.state.type} name="type"
								onChange={this.handleInputChange}
							>
								{this.sports.map(s => (
									<option key={s} value={s}>{s}</option>
								))}
								<option value="">Select Sport</option>
							</select>
						</div>

						<ErrMsgs className="help is-danger" for={this.state.errors.type}
							hiddenClass="is-hidden" 
						>
							<span for="required">This field is required</span>
						</ErrMsgs>
					</div>
				</div>
				<div className="field">
					<div className="control">
						<label className="radio">
							<input type="radio" name="isOpen" value="yes"
								checked={this.state.isOpen == true}
								onChange={this.handleRadioChange}
							/>
							Open
						</label>
						<label className="radio">
							<input type="radio" name="isOpen" value="no"
								checked={this.state.isOpen != true}
								onChange={this.handleRadioChange}
							/>
							Invitation-Only
						</label>

						<span className="help">
							Open leagues can be joined by all players
						</span>
					</div>
				</div>
				<div className="field">
					<div className="control">
						<label className="radio">
							<input type="radio" name="isPrivate" value="no"
								checked={this.state.isPrivate != true}
								onChange={this.handleRadioChange}
							/>
							Public
						</label>
						<label className="radio">
							<input type="radio" name="isPrivate" value="yes"
								checked={this.state.isPrivate == true}
								onChange={this.handleRadioChange}
							/>
							Private
						</label>

						<span className="help">
							Private leagues hide player roster from outsiders
						</span>
					</div>
				</div>
				<div className="field">
					<div className="control">
						<AsyncButton className="button"
							loadingClass="is-loading"
							text="Create League"
							onClick={this.submit}
						/>
					</div>
				</div>
			</form>
		);
	}
}

export default LeagueCreator;
