import React from "react";
import SlotMachine from "./slot-machine";

class Welcome extends React.Component {
	constructor(props) {
		super(props);

		this.slots = [
			"squash",
			"tennis",
			"soccer",
			"basketball",
			"DOTA",
			"tic-tac-toe",
			"laser tag",
			"Yu-Gi-Oh!",		
		];

		this.showJoinLeague = this.showJoinLeague.bind(this);
		this.showCreateLeague = this.showCreateLeague.bind(this);

		this.state = {
			isJoinShown: false,
			isCreateShown: false,
		};
	}

	showJoinLeague() {
		this.setState({
			isJoinShown: true,
			isCreateShown: false,
		});
		console.log("isJoinShown true");
	}

	showCreateLeague() {
		this.setState({
			isJoinShown: false,
			isCreateShown: true,
		});
	}

	render() {
		return (
		<div className="container">
			<div className="hero">
				<div className="hero-body">
					<div className="container">
						<h1 className="title">
							League-Hub
						</h1>
						<h2 className="subtitle" style={{display: "flex", alignItems: "center"}}>
							A hub for <SlotMachine opts={this.slots}>
							</SlotMachine> leagues
						</h2>
					</div>
				</div>
			</div>

			<div className="level">
				<div className="level-item has-text-centered">
					<div>
						<p className="title">12345</p>
						<p className="heading">Matches</p>
					</div>
				</div>
				<div className="level-item has-text-centered">
					<div>
						<p className="title">1234</p>
						<p className="heading">Players</p>
					</div>
				</div>
				<div className="level-item has-text-centered">
					<div>
						<p className="title">123</p>
						<p className="heading">Leagues</p>
					</div>
				</div>
			</div>

			<section className="section is-medium level">
				<p className="level-item">
					<a className={`level-item button subtitle ${this.state.isJoinShown ? "is-hidden" : ""}`} onClick={this.showJoinLeague}>
						Find a league to join
					</a>

					<span className={this.state.isJoinShown ? "" : "is-hidden"}>
						join
					</span>
				</p>
				<p className="level-item">
					<a className={`level-item button subtitle ${this.state.isCreateShown ? "is-hidden" : ""}`} onClick={this.showCreateLeague}>
						Create your own league
					</a>

					<span className={this.state.isCreateShown ? "" : "is-hidden"}>
						create
					</span>
				</p>
			</section>
		</div>
		)
	}
};

export default Welcome;
