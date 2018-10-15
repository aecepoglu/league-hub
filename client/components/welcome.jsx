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

			<section className="section is-medium has-text-centered">
				<a className="button subtitle">
					Find a league to join
				</a> <a className="button subtitle">
					or Create your own
				</a>
			</section>
		</div>
		)
	}
};

export default Welcome;
