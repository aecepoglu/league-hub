const fetch = require("isomorphic-fetch");

exports.Fetch = function(params) {
	return fetch(window.location.origin + "/graphql", {
		method: "POST",
		headers: {
			"content-type": "application/json"
		},
		body: JSON.stringify(params)
	}).then(resp => resp.json());
}

exports.Send = function(query, variables) {
	exports.Fetch({
		query: query,
		variables: variables
	}).then(x => {
		if (x.errors) {
			return Promise.reject(x.errors);
		}
		return x.data;
	});
};
