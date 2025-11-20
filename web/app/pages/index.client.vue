<template>
  <div>
    Welcome
    <button @click="handleUploadTest">Upload Test</button>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from "vue";
import { homePostSpeakClient, homePostUploadClient } from "@/clients/axios";
import type { HomePostUploadMutationRequest } from "~/clients/types";

// test backend call on mounted
console.log("Calling backend before mount API...");
onBeforeMount(async () => {
	console.log("Calling backend API...");
	const response = await homePostSpeakClient({
		message: "Hello from Nuxt 3 frontend!",
	});
	console.log("Response from backend:", response);
});

const handleUploadTest = async () => {
	console.log("Starting upload test...");
	const requestData: HomePostUploadMutationRequest = {};
	const fileContent = new Blob(["This is a test file content"], {
		type: "text/plain",
	});
	requestData["file"] = new File([fileContent], "testfile.txt");
	requestData["bulk-files"] = [];
	for (let i = 0; i < 3; i++) {
		const anotherFileContent = new Blob(
			[`This is test file content number ${i + 1}`],
			{ type: "text/plain" },
		);
		requestData["bulk-files"].push(new File([anotherFileContent], `file${i + 1}.txt`));
	}

	try {
		const response = await homePostUploadClient(requestData);

		const data = await response.data;
		console.log("Upload response:", data);
	} catch (error) {
		console.error("Error during file upload:", error);
	}
};
</script>