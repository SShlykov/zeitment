module.exports = {
  extends: [
    "eslint:recommended",
    "plugin:vue/vue3-recommended",
  ],
  rules: {
    "no-multiple-empty-lines": ["error", {max: 2, "maxBOF": 0}],
    "indent": ["error", 2],
    "no-mixed-spaces-and-tabs": "error",
    "vue/attribute-hyphenation": "off",
    "vue/html-self-closing": ["error", {
      "html": {
        "void": "never",
        "normal": "always",
        "component": "always"
      },
      "svg": "always",
      "math": "always"
    }],
    "vue/no-reserved-component-names": "off",
    "vue/html-indent": ["error", 2, {
      "attribute": 1,
      "baseIndent": 1,
      "closeBracket": 0,
      "alignAttributesVertically": true,
      "ignores": []
    }]
  }
}
