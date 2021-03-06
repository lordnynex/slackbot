// Generated by ego.
// DO NOT EDIT

package command
import (
"fmt"
"html"
"io"
)
//line command/templates/plugin_template.go.ego:1
 func PluginTmpl(w io.Writer, pkgName string) error  {
//line command/templates/plugin_template.go.ego:1
_, _ = fmt.Fprint(w, "package ")
//line command/templates/plugin_template.go.ego:1
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  pkgName )))
//line command/templates/plugin_template.go.ego:2
_, _ = fmt.Fprint(w, "\n\nimport (\n\t\"github.com/kyokomi/slackbot/plugins\"\n)\n\ntype plugin struct {\n}\n\nfunc NewPlugin() plugins.BotMessagePlugin {\n\treturn &plugin{}\n}\n\nfunc (p *plugin) CheckMessage(_ plugins.BotEvent, message string) (bool, string) {\n\t// TODO: execute message check\n\treturn true, message\n}\n\nfunc (p *plugin) DoAction(event plugins.BotEvent, message string) bool {\n\t// TODO: reply message action\n\tevent.Reply(message)\n\treturn true // next ok\n}\n\nvar _ plugins.BotMessagePlugin = (*plugin)(nil)\n")
return nil
}
//line command/templates/plugin_test_template.go.ego:1
 func PluginTestTmpl(w io.Writer, pkgName string) error  {
//line command/templates/plugin_test_template.go.ego:1
_, _ = fmt.Fprint(w, "package ")
//line command/templates/plugin_test_template.go.ego:1
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  pkgName )))
//line command/templates/plugin_test_template.go.ego:1
_, _ = fmt.Fprint(w, "_test\n\nimport (\n\t\"testing\"\n\n\t\"github.com/kyokomi/slackbot/plugins\"\n\t\"github.com/kyokomi/slackbot/plugins/")
//line command/templates/plugin_test_template.go.ego:7
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  pkgName )))
//line command/templates/plugin_test_template.go.ego:7
_, _ = fmt.Fprint(w, "\"\n)\n\nvar testEvent = plugins.NewTestEvent(\"test\")\n\nfunc TestCheckMessage(t *testing.T) {\n\tp := ")
//line command/templates/plugin_test_template.go.ego:13
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  pkgName )))
//line command/templates/plugin_test_template.go.ego:13
_, _ = fmt.Fprint(w, ".NewPlugin()\n\tok, _ := p.CheckMessage(testEvent, testEvent.BaseText())\n\tif !ok {\n\t\tt.Errorf(\"ERROR check = NG\")\n\t}\n}\n\nfunc TestDoAction(t *testing.T) {\n\tp := ")
//line command/templates/plugin_test_template.go.ego:21
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  pkgName )))
//line command/templates/plugin_test_template.go.ego:21
_, _ = fmt.Fprint(w, ".NewPlugin()\n\n\tnext := p.DoAction(testEvent, testEvent.BaseText())\n\n\tif next != true {\n\t\tt.Errorf(\"ERROR next != true\")\n\t}\n}\n")
return nil
}
