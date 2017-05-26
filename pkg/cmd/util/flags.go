package util

import (
	"github.com/spf13/cobra"
	kapi "k8s.io/kubernetes/pkg/api"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/resource"
)

func AddGetFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("namespace", "n", kapi.NamespaceDefault, "List the requested object(s) from this namespace.")
	cmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.")
	cmd.Flags().Bool("all-namespaces", false, "If present, list the requested object(s) across all namespaces. Namespace specified with --namespace will be ignored.")
	cmd.Flags().Bool("show-kind", false, "If present, list the resource type for the requested object(s).")
	cmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|wide|name.")
	cmd.Flags().BoolP("show-all", "a", false, "When printing, show all resources (default hide terminated pods.)")
	cmd.Flags().Bool("show-labels", false, "When printing, show all labels as the last column (default hide labels column)")
}

func AddCreateFlags(cmd *cobra.Command, options *resource.FilenameOptions) {
	cmd.Flags().StringP("namespace", "n", kapi.NamespaceDefault, "Create object(s) in this namespace.")
	usage := "create the resource"
	AddFilenameOptionFlags(cmd, options, usage)
}

func AddDeleteFlags(cmd *cobra.Command, options *resource.FilenameOptions) {
	cmd.Flags().StringP("namespace", "n", kapi.NamespaceDefault, "Delete object(s) from this namespace.")
	cmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on.")
	cmd.Flags().StringP("output", "o", "", "Output mode. Use \"-o name\" for shorter output (resource/name).")
	usage := "delete the resource"
	AddFilenameOptionFlags(cmd, options, usage)
}

func AddDescribeFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("namespace", "n", kapi.NamespaceDefault, "Describe object(s) from this namespace.")
	cmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.")
	cmd.Flags().Bool("all-namespaces", false, "If present, describe the requested object(s) across all namespaces. Namespace specified with --namespace will be ignored.")
}

func AddEditFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("namespace", "n", kapi.NamespaceDefault, "Edit object(s) in this namespace.")
	cmd.Flags().StringP("output", "o", "yaml", "Output format. One of: yaml|json.")
}

func AddInitFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("namespace", "n", kapi.NamespaceDefault, "Namespace name. Operator will be deployed in this namespace.")
	cmd.Flags().String("version", "canary", "Operator version")
	cmd.Flags().Bool("upgrade", false, "If present, Upgrade operator to use provided version")
}

func AddFilenameOptionFlags(cmd *cobra.Command, options *resource.FilenameOptions, usage string) {
	cmd.Flags().StringSliceVarP(&options.Filenames, "filename", "f", options.Filenames, "Filename to use to "+usage)
	cmd.Flags().BoolVarP(&options.Recursive, "recursive", "R", options.Recursive, "Process the directory used in -f, --filename recursively.")
}

func GetNamespace(cmd *cobra.Command) (string, bool) {
	return cmdutil.GetFlagString(cmd, "namespace"), cmd.Flags().Changed("namespace")
}