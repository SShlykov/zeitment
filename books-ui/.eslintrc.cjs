module.exports = {
  extends: [
    "plugin:vue/vue3-recommended",
  ],
  rules: {
    "no-mixed-spaces-and-tabs": "error",
    "vue/multi-word-component-names": ["error", {
      "ignores": [
        "Head",
        "Standardization",
        "Metrology",
        "Certification",
        "Loader",
        "Tabs",
        "Table",
        "Toast",
      ]
    }],
    "vue/multi-word-component-names": "off",
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
    "vue/no-use-v-if-with-v-for": "off"
  }
}
