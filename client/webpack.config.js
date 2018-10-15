const path = require("path");
//
//console.log(path.join(__dirname, "index.jsx"));

module.exports = {
	devtool: "eval-source-map",
	mode: process.env["WEBPACK_MODE"] || "development",

	entry: "./index.jsx",
	output: {
		path: path.resolve("./public"),
		filename: "bundle.js"
	},

	resolve: {
		extensions: [".js", ".jsx"]
	},
	module: {
		rules: [{
			test: /\.jsx?$/,
			exclude: /node_modules/,
			loader: "babel-loader",
		}, {
			test: /\.flow$/,
			loader: "ignore-loader",
		}, {
			test: /\.css$/,
			use: [
				"style-loader",
				{
					loader: "css-loader",
					options: {
						modules: true,
						localIdentName: "[sha1:hash:hex:4]",
					}
				}
			]
		}]
	}
};
