import { check, sleep } from 'k6';
import http from 'k6/http';

export default function () {
    const res = http.post(
        `http://${__ENV.HOSTNAME}:4001/api/event`, 
        JSON.stringify(
            {
                "id": "6d4964bb-5be3-421c-85f9-75414417bb3a",
                "name": "My Event",
                "date": "2023-04-20T14:00:00Z",
                "languages": [
                    "English",
                    "French",
                    "Lithuanian"
                ],
                "videoQuality": [
                    "720p",
                    "1080p",
                    "1440p",
                    "2160p"
                ],
                "audioQuality": [
                    "High",
                    "Medium",
                    "Low"
                ],
                "invitees": [
                    "example1@gmail.com",
                    "example2@gmail.com",
                    "example3@gmail.com",
                    "example4@gmail.com",
                    "example5@gmail.com",
                    "example6@gmail.com",
                    "example7@gmail.com",
                    "example8@gmail.com",
                    "example9@gmail.com"
                ],
                "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged."
            }
        ),
        {
            headers: { 'Content-Type': 'application/json' },
        }
    );
    check(res, { 'status was 201': (r) => r.status == 201 });
    sleep(1);
}