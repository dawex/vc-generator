/**
 * Documentation : https://redocly.com/docs/cli/custom-plugins/custom-decorators
 */
const OpenAPISnippet = require('openapi-snippet');

module.exports = function CodeSnippets() {
	return {
		id: 'code-snippets',
		decorators: {
			oas3: {
				'add': AddCodeSnippet,
			},
		},
	};
};

/**
 * Add a code snippets on operations
 *
 * Inspired by https://github.com/cdwv/oas3-api-snippet-enricher
 @param targets The
 */
function AddCodeSnippet({targets}) {
	return {
		Root: {
			leave(schema) {
				const snippets = OpenAPISnippet.getSnippets(schema, targets);

				snippets.filter(snippet => /get|put|post|delete|patch|options|head|trace/i.test(snippet.method)).forEach(snippet => {
					const path = Object.keys(schema.paths).find(p => snippet.url.endsWith(p));
					const method = snippet.method.toLocaleLowerCase();
					if (path) {
						schema.paths[path][method]["x-codeSamples"] = snippet.snippets.map(sample => ({ "lang": sample.title, "source": sample.content }));
					}
				})
			}
		},
	}
}
