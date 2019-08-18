package grep

import (
	"regexp"
)

func init() {
	global.fileTypesMap = map[string]FileType{
		"go": {
			Patterns: []string{"*.go"},
		},
		"cc": {
			Patterns: []string{"*.c", "*.h", "*.xs"},
		},
		"cpp": {
			Patterns: []string{"*.cpp", "*.cc", "*.cxx", "*.m", "*.hpp", "*.hh", "*.h", "*.hxx"},
		},
		"html": {
			Patterns: []string{"*.htm", "*.html", "*.shtml", "*.xhtml"},
		},
		"groovy": {
			Patterns: []string{"*.groovy", "*.gtmpl", "*.gpp", "*.grunit", "*.gradle"},
		},
		"java": {
			Patterns: []string{"*.java", "*.properties"},
		},
		"jsp": {
			Patterns: []string{"*.jsp", "*.jspx", "*.jhtm", "*.jhtml"},
		},
		"perl": {
			Patterns:     []string{"*.pl", "*.pm", "*.pod", "*.t"},
			ShebangRegex: regexp.MustCompile(`^#!.*\bperl\b`),
		},
		"php": {
			Patterns:     []string{"*.php", "*.phpt", "*.php3", "*.php4", "*.php5", "*.phtml"},
			ShebangRegex: regexp.MustCompile(`^#!.*\bphp\b`),
		},
		"ruby": {
			Patterns:     []string{"*.rb", "*.rhtml", "*.rjs", "*.rxml", "*.erb", "*.rake", "*.spec", "Rakefile"},
			ShebangRegex: regexp.MustCompile(`^#!.*\bruby\b`),
		},
		"python": {
			Patterns:     []string{"*.py", "*.pyw", "*.pyx", "SConstruct"},
			ShebangRegex: regexp.MustCompile(`^#!.*\bpython[0-9.]*\b`),
		},
		"shell": {
			Patterns:     []string{"*.sh", "*.bash", "*.csh", "*.tcsh", "*.ksh", "*.zsh"},
			ShebangRegex: regexp.MustCompile(`^#!.*\b(?:ba|t?c|k|z)?sh\b`),
		},
		"xml": {
			Patterns:     []string{"*.xml", "*.dtd", "*.xsl", "*.xslt", "*.ent"},
			ShebangRegex: regexp.MustCompile(`<\?xml`),
		},
	}
}
