// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Register() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form id=\"register-card\" class=\"mx-auto max-w-md\" autocomplete=\"on\" hx-post=\"/auth/register\" hx-target=\"main\" hx-swap=\"transition:true\" hx-push-url=\"/login\"><div class=\"flex flex-col p-6 space-y-1 text-center\"><h2 class=\"whitespace-nowrap tracking-tight text-2xl font-bold\">Create an account</h2><p class=\"text-sm text-muted-foreground\">Get started with our services by creating a new account.</p></div><div class=\"p-6 space-y-4\"><div class=\"grid gap-2\"><label class=\"text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70\" for=\"name\">Full Name</label> <input class=\"flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50\" id=\"name\" placeholder=\"John Doe\" autocomplete=\"on\"></div><div class=\"grid gap-2\"><label class=\"text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70\" for=\"phone\">Phone Number</label> <input class=\"flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50\" id=\"phone\" placeholder=\"+1 (555) 555-5555\" autocomplete=\"on\"></div><div class=\"grid gap-2\"><label class=\"text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70\" for=\"email\">Email</label> <input class=\"flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50\" id=\"email\" placeholder=\"example@email.com\" type=\"email\" autocomplete=\"on\"></div><div class=\"grid gap-2\"><label class=\"text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70\" for=\"password\">Password</label> <input class=\"flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50\" id=\"password\" type=\"password\"></div><div class=\"grid gap-2\"><label class=\"text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70\" for=\"confirm-password\">Confirm Password</label> <input class=\"flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50\" id=\"confirm-password\" type=\"password\"></div></div><div class=\"flex justify-center\"><p class=\"text-sm text-muted-foreground\">Already have an account?  <a href=\"#\" hx-get=\"/login\" hx-target=\"main\" hx-swap=\"transition:true\" hx-push-url=\"/login\" class=\"text-sm text-primary-foreground hover:text-accent hover:underline\"><b>Sign in</b></a></p></div><div class=\"items-center p-6 flex justify-end gap-2\"><button hx-get=\"/\" hx-target=\"main\" hx-swap=\"transition:true\" hx-push-url=\"/\" class=\"inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2\">Cancel</button> <button class=\"inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2\">Sign Up</button></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Login() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form id=\"login-card\" class=\"mx-auto max-w-md\" autocomplete=\"on\" hx-post=\"/auth/login\" hx-target=\"main\" hx-swap=\"transition:true\" hx-push-url=\"/\"><div class=\"flex flex-col p-6 space-y-1 text-center\"><h2 class=\"whitespace-nowrap tracking-tight text-2xl font-bold\">Sign in to your account</h2><p class=\"text-sm text-muted-foreground\">Enter your credentials to access your account.</p></div><div class=\"p-6 space-y-4\"><div class=\"grid gap-2\"><label class=\"text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70\" for=\"email\">Email</label> <input class=\"flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50\" id=\"email\" placeholder=\"example@email.com\" type=\"email\" autocomplete=\"on\"></div><div class=\"grid gap-2\"><label class=\"text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70\" for=\"password\">Password</label> <input class=\"flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50\" id=\"password\" type=\"password\"></div></div><div class=\"flex justify-center\"><p class=\"text-sm text-muted-foreground\">Don't have an account?  <a href=\"#\" hx-get=\"/register\" hx-target=\"main\" hx-swap=\"transition:true\" hx-push-url=\"/register\" class=\"text-sm text-primary-foreground hover:text-accent hover:underline\"><b>Sign up</b></a></p></div><div class=\"items-center p-6 flex justify-end gap-2\"><button hx-get=\"/\" hx-target=\"main\" hx-swap=\"transition:true\" hx-push-url=\"/\" class=\"inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2\">Cancel</button> <button class=\"inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2\">Sign In</button></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
