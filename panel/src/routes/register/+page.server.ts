import type { Actions } from './$types';
import { redirect } from "@sveltejs/kit"

export const actions = {
	default: async (event) => {
        console.log("Register action triggered");
        const formData = await event.request.formData();
        const email = formData.get('email') as string;
        const password = formData.get('password') as string;
        const confirmPassword = formData.get('confirm_password') as string;

        const request = await fetch("http://127.0.0.1:8080/api/user/signup", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                email: email,
                password: password,
                confirm_password: confirmPassword
            })
        })
        if (request.status === 200) {
            const data = await request.json();
            event.locals.token = data.token;
            event.cookies.set("token", data.token, {
                path: "/"
            })
            return redirect(302, "/");
        }
	}
} satisfies Actions;