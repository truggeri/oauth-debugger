const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const SveltePreProcess = require("svelte-preprocess");
const path = require("path");

const mode = process.env.NODE_ENV || "development";
const prod = mode === "production";

module.exports = {
	entry: {
		bundle: ["./src/main.js"]
	},
	resolve: {
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
	mode,
	plugins: [
		new MiniCssExtractPlugin({
			filename: "[name].css"
		})
	],
	devtool: prod ? false: "source-map"
};
