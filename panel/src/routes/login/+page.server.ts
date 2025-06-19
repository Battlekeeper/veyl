import type { Actions } from './$types';
import { redirect } from "@sveltejs/kit"

export const actions = {
    login: async (event) => {
        const formData = await event.request.formData();
        const email = formData.get('email') as string;
        const password = formData.get('password') as string;

        const request = await fetch("http://127.0.0.1:8080/api/user/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                email: email,
                password: password,
            })
        })
        if (request.status === 200) {
            const data = await request.json();
            event.locals.token = data.token;
            event.cookies.set("token", data.token, {
                path: "/"
            })
            return redirect(302, "/panel");
        }
    }
} satisfies Actions;