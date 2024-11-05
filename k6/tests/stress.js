import http from 'k6/http';
import { check, group, sleep } from 'k6';

export let options = {
  stages: [
    { duration: '10s', target: 1000 },  // Ramp-up to 1000 users
    { duration: '60s', target: 10000 },  // Stay at 5000 users
    { duration: '10s', target: 0 },   // Ramp-down to 0 users
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'], // 95 درصد درخواست‌ها باید در کمتر از 500 میلی‌ثانیه پاسخ داده شوند
    http_req_failed: ['rate<0.05'], // نرخ شکست‌ها کمتر از 5% باید باشد.
  },
};

export default function () {
  group('API Home', () => {
    let res1 = http.get('http://localhost:8080/');
    check(res1, { 'status is 200': (r) => r.status === 200 });
  });

  group('API Users', () => {
    let res2 = http.get('http://localhost:8080/users');
    check(res2, { 'status is 200': (r) => r.status === 200 });
  });
}