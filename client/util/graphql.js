const fetch = require("isomorphic-fetch");

function request(params) {
	return fetch(window.location.origin + "/graphql", {
		method: "POST",
		headers: {
			"content-type": "application/json"
		},
		body: JSON.stringify(params)
	}).then(resp => resp.json());
}

function send(query, variables) {
	return request({
		query: query,
		variables: variables
	}).then(x => {
		if (x.errors) {
			return Promise.reject(x.errors);
		}
		return x.data;
	});
}

export {request, send};
