import React from "react";
import style from "./style.css";

class SlotMachine extends React.Component {
	constructor(props) {
		super(props);
		this.options = [].concat(
			props.opts,
			props.opts,
			props.opts
		).map(x => ({val: x}));;
		this.state = {y: 0};
		this.timer = setInterval(this.tick.bind(this), 1500);
	}

	componentWillUnmount() {
		super.componentWillUnmount();
		clearInterval(this.timer);
	}

	tick() {
		let i = Math.floor(Math.random() * this.options.length);
		let y = this.options[i].ref.getBoundingClientRect().y;
		let y0 = this.ref.getBoundingClientRect().y;
		this.setState({y: y0 - y});
	}

	render() {
		return (<span className={style.container} style={{height: "1.25em"}}>
			<div style={{top: this.state.y}} ref={x => this.ref = x}>
				{this.options.map((o, i) => (
					<div key={i} ref={x => o.ref = x}>{o.val}</div>
				))}
			</div>
		</span>);
	}
}

export default SlotMachine;
