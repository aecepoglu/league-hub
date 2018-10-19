import React from "react";

function ErrMsgs(props) {
	return <p className={`${props.className} ${props.for ? "" : props.hiddenClass}`}>
		{props.for && React.Children.map(props.children, (c, i) => {
			return props.for[c.props.for] && <span key={i}>{c.props.children}</span>
		})}
	</p>;
}

export default ErrMsgs;
