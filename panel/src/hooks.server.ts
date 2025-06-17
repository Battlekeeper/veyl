import type { Handle } from "@sveltejs/kit"
import { redirect } from "@sveltejs/kit"


export const handle: Handle = async ({ event, resolve }) => {
    const token = event.cookies.get("token")

    if (!token) {
        if (!event.url.pathname.includes("/login") && !event.url.pathname.includes("/register") && !event.url.pathname.includes("/logout") && event.url.pathname !== "/") {
            event.cookies.delete("token", {
                path: '/'
            })
            return redirect(302, `/logout`)
        }
        return resolve(event)
    }

    const response = await fetch("http://127.0.0.1:8080/api/user", {
        headers: {
            Authorization: "Bearer " + token,
        },
    })


    if (response.status !== 200) {
        if (event.url.pathname !== "/login") {
            event.cookies.delete("token", {
                path: '/'
            })
            return redirect(302, `/logout`)
        }
    }

    const data = await response.json()
    event.locals.user = data
    event.locals.token = token

    return resolve(event)
}
