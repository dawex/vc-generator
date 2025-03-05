<template>
	<div className="flex flex-col min-h-screen overflow-hidden">
		<!-- Site header -->
		<Header />
		<!-- Page content -->
		<main class="grow">
			<section class="relative">
				<div>
					<!-- Main content -->
					<div>
						<div>
							<div class="top-0 left-0 md:top-20 outer-container bg-white">
								<div ref="redoc" />
							</div>
						</div>
					</div>
				</div>
			</section>
		</main>
	</div>
</template>

<script>
import { ref } from 'vue'

import Header from '/src/partials/Header.vue'

export default {
	name: 'Redoc',
	components: {
		Header,
	},
	setup() {
		const sidebarOpen = ref(false)
		const redoc = ref(null);

		return {
			sidebarOpen,
			redoc,
		}
	},
	mounted() {
		this.initRedoc('vc-generator.yaml') // force only one file vc-generator.yaml
	},
	methods: {
		initRedoc(filename) {
			window.Redoc.init(
				__APP_BASE_URL__ + filename,
				{
					hideLogo: true,
					disableSearch: true
				},
				this.redoc
			)
		}
	}
}
</script>

<style>
.outer-container {
	display: flex;
	flex-direction: column;
	position: fixed;
	bottom: 0;
	right: 0;
	scroll-padding: 0 !important;
	scroll-behavior: unset !important;
	overflow-y: scroll;
}

.bg-white {
	background-color: aliceblue;
}
</style>
