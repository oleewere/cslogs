package grep

import (
	"regexp"
)

func init() {
	global.fileTypesMap = map[string]FileType{
		"go": FileType{
			Patterns: []string{"*.go"},
		},
		"cc": FileType{
			Patterns: []string{"*.c", "*.h", "*.xs"},
		},
		"cpp": FileType{
			Patterns: []string{"*.cpp", "*.cc", "*.cxx", "*.m", "*.hpp", "*.hh", "*.h", "*.hxx"},
		},
		"html": FileType{
			Patterns: []string{"*.htm", "*.html", "*.shtml", "*.xhtml"},
		},
		"groovy": FileType{
			Patterns: []string{"*.groovy", "*.gtmpl", "*.gpp", "*.grunit", "*.gradle"},
		},
		"java": FileType{
			Patterns: []string{"*.java", "*.properties"},
		},
		"jsp": FileType{
			Patterns: []string{"*.jsp", "*.jspx", "*.jhtm", "*.jhtml"},
		},
		"perl": FileType{
			Patterns:     []string{"*.pl", "*.pm", "*.pod", "*.t"},
			ShebangRegex: regexp.MustCompile(`^#!.*\bperl\b`),
		},
		"php": FileType{
			Patterns:     []string{"*.php", "*.phpt", "*.php3", "*.php4", "*.php5", "*.phtml"},
			ShebangRegex: regexp.MustCompile(`^#!.*\bphp\b`),
		},
		"ruby": FileType{
			Patterns:     []string{"*.rb", "*.rhtml", "*.rjs", "*.rxml", "*.erb", "*.rake", "*.spec", "Rakefile"},
			ShebangRegex: regexp.MustCompile(`^#!.*\bruby\b`),
		},
		"python": FileType{
			Patterns:     []string{"*.py", "*.pyw", "*.pyx", "SConstruct"},
			ShebangRegex: regexp.MustCompile(`^#!.*\bpython[0-9.]*\b`),
		},
		"shell": FileType{
			Patterns:     []string{"*.sh", "*.bash", "*.csh", "*.tcsh", "*.ksh", "*.zsh"},
			ShebangRegex: regexp.MustCompile(`^#!.*\b(?:ba|t?c|k|z)?sh\b`),
		},
		"xml": FileType{
			Patterns:     []string{"*.xml", "*.dtd", "*.xsl", "*.xslt", "*.ent"},
			ShebangRegex: regexp.MustCompile(`<\?xml`),
		},
	}
}
