#!/usr/bin/env node

const { exec, execSync } = require("child_process");

function sshToHttps(sshUrl) {
  return sshUrl
    .replace(/^git@github\.com:/, "https://github.com/")
    .replace(/\.git$/, "");
}

try {
  const branchName = execSync("git branch --show-current").toString().trim();


  exec("git remote get-url origin", (err, stdout, stderr) => {
    if (err || stderr) {
      console.error(err || stderr);
    }

    const url = `${sshToHttps(stdout.trim())}/pull/${branchName}`;

    const command =
      process.platform === "win32" ? `start ${url}` : `open ${url}`;
    exec(command, (err) => {
      if (err) {
        console.error("Failed to open URL:", err);
      } else {
        console.log("Opening URL:", url);
      }
    });
  });
} catch (error) {}
