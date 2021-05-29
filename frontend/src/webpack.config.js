const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const SveltePreProcess = require("svelte-preprocess");
const path = require("path");
const webpack = require("webpack");

const mode = process.env.NODE_ENV || "development";
const prod = mode === "production";

const define = new webpack.DefinePlugin({
	PRODUCTION: JSON.stringify(prod),
});
const provide = new webpack.ProvidePlugin({
	CONFIG: path.resolve(__dirname, prod ? "config/production.js" : "config/dev.js")
});

const svelteConfig = {
	resolver: {
		alias: {
			svelte: path.resolve("node_modules", "svelte")
		},
		extensions: [".tsx", ".ts", ".mjs", ".js", ".svelte"],
		mainFields: ["svelte", "browser", "module", "main"]
	},
	output: {
		path: __dirname + "/public",
		filename: "[name].js",
		chunkFilename: "[name].[id].js"
	},
	module: {
		rules: [
			{
				test: /\.svelte$/,
				use: {
					loader: "svelte-loader",
					options: {
						emitCss: true,
						hotReload: true,
						preprocess: SveltePreProcess({}),
					}
				}
			},
			{
				test: /\.css$/,
				use: [
					/**
					 * MiniCssExtractPlugin doesn't support HMR.
					 * For developing, use "style-loader" instead.
					 * */
					prod ? MiniCssExtractPlugin.loader : "style-loader",
					"css-loader"
				]
			},
			{
				test: /\.tsx?$/,
				use: "ts-loader",
				exclude: /node_modules/,
			}
		]
	},
	plugins: [
		new MiniCssExtractPlugin({
			filename: "[name].css"
		}),
		provide,
		define
	],
	devtool: prod ? false: "source-map",
}

module.exports = [
	{
		entry: {
			login: ["./login.js"],
			setup: ["./setup.js"],
		},
		resolve: svelteConfig.resolver,
		output: svelteConfig.output,
		module: svelteConfig.module,
		mode,
		plugins: svelteConfig.plugins,
		devtool: svelteConfig.devtool,
	},
];
