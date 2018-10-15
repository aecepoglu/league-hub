import React from "react";

class SlotMachine extends React.Component {
	constructor(props) {
		super(props);
		this.options = props.opts;
	}

	render() {
		return (<span id="abcd" className="slot-machine" style={{height: "1.25em"}}>
			<div>
				{this.options.map((o, i) => (
					<div key={o}>{o}</div>
				))}
			</div>
		</span>);
	}
}

class Welcome extends React.Component {
	constructor(props) {
		super(props);

		this.timer = setInterval(() => {
		}, 2500);

		this.slots = [
			"squash",
			"tennis",
			"soccer",
			"basketball",
			"DOTA",
			"underground fightclub",
			"laser tag",
		];
	}

	render() {
		return (
		<section className="hero">
			<div className="hero-body">
				<div className="container">
					<h1 className="title">
						League-Hub
					</h1>
					<h2 className="subtitle">
						Create your own <SlotMachine id="abc" opts={this.slots}>
						</SlotMachine> league
					</h2>
				</div>
			</div>
		</section>
		)
	}
};

export default Welcome;
