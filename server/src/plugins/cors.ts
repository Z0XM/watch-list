import fp from 'fastify-plugin';
import cors from '@fastify/cors';

export default fp(async (fastify) => {
	fastify.register(cors, {
		origin: ['*'],
		methods: ['GET', 'PUT', 'PATCH', 'POST', 'DELETE']
	});
});
