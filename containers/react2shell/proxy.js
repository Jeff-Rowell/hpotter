import { NextResponse } from "next/server";
import { NextRequest } from "next/server";
import { headers } from "next/headers";

export async function proxy(request) {
    process.stdout.write("REQUEST:");

    const basicAuth = request.headers.get("authorization");
    if (basicAuth && basicAuth.startsWith("Basic ")) {
        const base64Credentials = basicAuth.split(" ")[1];
        const decodedCredentials = atob(base64Credentials);
        const [username, password] = decodedCredentials.split(":");

        if (username.length > 0 && password.length > 0) {
            const usernameBodyLiteral = username.replace(/(\r\n|\n|\r)/g, '')
            process.stdout.write(" Username=" + usernameBodyLiteral);
            const passwordBodyLiteral = password.replace(/(\r\n|\n|\r)/g, '')
            process.stdout.write(", Password=" + passwordBodyLiteral + ",");
        }
    }

    if (request.method == "POST") {
        const requestBody = await request.clone().text();
        const requestBodyLiteral = requestBody.replace(/(\r\n|\n|\r)/g, '\\n')
        process.stdout.write(" Payload=" + requestBodyLiteral);
    } else {
        process.stdout.write(" Payload={}");
    }

    process.stdout.write(", Method=" + request.method);
    process.stdout.write(", URI=" + request.url);
    process.stdout.write(", Headers=" + JSON.stringify(Object.fromEntries(request.headers)));
    process.stdout.write("\n");

    const nextAction = request.headers.get("next-action");
    const rscActionId = request.headers.get("rsc-action-id");
    if (nextAction || rscActionId) {
        return new NextResponse('{"status":"success"}', { status: 200 });
    }
}

export const config = {
    matcher: '/:path*'
};
