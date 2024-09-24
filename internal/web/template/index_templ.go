// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package template

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func LandingPage() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"hero min-h-screen\" style=\"background-image: url(/assets/img/placeholder.svg);\"><div class=\"hero-overlay bg-opacity-60\"></div><div class=\"hero-content text-neutral-content text-center\"><div class=\"max-w-md\"><h1 class=\"mb-5 text-5xl font-bold drop-shadow\">Elevate Your Beauty at SEA Salon</h1><p class=\"mb-5 drop-shadow\">Experience the finest hair and beauty services in town. Our talented team of stylists and aestheticians are dedicated to helping you look and feel your best.</p><a href=\"#services\" class=\"btn btn-primary\">Our Services</a> <a href=\"#contact\" class=\"btn btn-secondary\">Contact</a></div></div></section><section id=\"services\" class=\"bg-base-100 py-12 md:py-24\"><div class=\"container px-4 mx-auto md:px-6\"><div class=\"flex flex-col items-center justify-center space-y-4 text-center\"><div class=\"space-y-2\"><h2 class=\"text-3xl font-bold tracking-tighter sm:text-4xl\">Our Services</h2><p class=\"max-w-[900px] md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed\">From haircuts and color to manicures and facials, we offer a wide range of services to help you look and feel your best.</p></div></div><div class=\"mx-auto grid max-w-5xl grid-cols-1 gap-6 py-12 sm:grid-cols-2 md:grid-cols-3 lg:gap-12\"><div class=\"flex flex-col items-center justify-center space-y-4 rounded-lg bg-inherit p-6 shadow-sm transition-colors hover:bg-accent\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"h-12 w-12 text-primary\"><circle cx=\"6\" cy=\"6\" r=\"3\"></circle> <path d=\"M8.12 8.12 12 12\"></path> <path d=\"M20 4 8.12 15.88\"></path> <circle cx=\"6\" cy=\"18\" r=\"3\"></circle> <path d=\"M14.8 14.8 20 20\"></path></svg><h3 class=\"text-xl font-bold\">Haircuts &amp; Styling</h3><p class=\"text-center\">Our skilled stylists will give you a fresh, modern look.</p></div><div class=\"flex flex-col items-center justify-center space-y-4 rounded-lg bg-inherit p-6 shadow-sm transition-colors hover:bg-accent\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"h-12 w-12 text-primary\"><path d=\"M2 13a6 6 0 1 0 12 0 4 4 0 1 0-8 0 2 2 0 0 0 4 0\"></path> <circle cx=\"10\" cy=\"13\" r=\"8\"></circle> <path d=\"M2 21h12c4.4 0 8-3.6 8-8V7a2 2 0 1 0-4 0v6\"></path> <path d=\"M18 3 19.1 5.2\"></path> <path d=\"M22 3 20.9 5.2\"></path></svg><h3 class=\"text-xl font-bold\">Manicures &amp; Pedicures</h3><p class=\"text-center\">Treat your hands and feet to a relaxing spa experience.</p></div><div class=\"flex flex-col items-center justify-center space-y-4 rounded-lg bg-inherit p-6 shadow-sm transition-colors hover:bg-accent\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"h-12 w-12 text-primary\"><path d=\"M3 7V5a2 2 0 0 1 2-2h2\"></path> <path d=\"M17 3h2a2 2 0 0 1 2 2v2\"></path> <path d=\"M21 17v2a2 2 0 0 1-2 2h-2\"></path> <path d=\"M7 21H5a2 2 0 0 1-2-2v-2\"></path> <path d=\"M8 14s1.5 2 4 2 4-2 4-2\"></path> <path d=\"M9 9h.01\"></path> <path d=\"M15 9h.01\"></path></svg><h3 class=\"text-xl font-bold\">Facial Treatments</h3><p class=\"text-center\">Rejuvenate your skin with our customized facial treatments.</p></div></div></div></section><section id=\"contact\" class=\"bg-base-200 py-12 md:py-24\"><div class=\"container px-4 mx-auto md:px-6\"><div class=\"grid gap-6 md:grid-cols-2 lg:gap-12\"><div class=\"space-y-4\"><h2 class=\"text-3xl font-bold tracking-tighter sm:text-4xl\">Get in Touch</h2><p class=\"md:text-xl\">We'd love to hear from you! Contact us to schedule an appointment or learn more about our services.</p><div class=\"grid gap-4\"><div class=\"flex items-center gap-2\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"h-6 w-6 text-primary\"><path d=\"M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z\"></path></svg> <a href=\"tel:+628123456789\" class=\"hover:text-accent\" rel=\"ugc\">+62 812-3456-789 (Thomas)</a></div><div class=\"flex items-center gap-2\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"h-6 w-6 text-primary\"><path d=\"M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z\"></path></svg> <a href=\"tel:+628164829372\" class=\"hover:text-accent\" rel=\"ugc\">+62 816-4829-372 (Sekar)</a></div></div></div><div class=\"space-y-4\"><h2 class=\"text-3xl font-bold tracking-tighter sm:text-4xl\">Visit Us</h2><p class=\"md:text-xl\">Our salon is conveniently located in the heart of the city. Stop by and experience the SEA Salon difference.</p><img src=\"/assets/img/placeholder.svg\" width=\"550\" height=\"310\" alt=\"Salon Location\" class=\"mx-auto aspect-video overflow-hidden rounded-xl object-cover object-center sm:w-full\"></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
