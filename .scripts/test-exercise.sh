#!/bin/bash

lab=$1
echo "Test instructions for each lab"
echo ""

# Set print command
PS4='\033[34m>> \033[0m'

for dir in $lab; do
	if [ -d "$dir" ]; then
		echo "#!/bin/bash" >"$dir/install.sh"
		echo "source .scripts/test-utils.sh" >>"$dir/install.sh"
		awk '/lab-instruction/ {found=1} found' "$dir/README.md" | sed -n '/```/,/```/p' | sed '/```/d' >>$dir/install.sh
		echo "----------------------------------------------------------------------------------------------------"
		cat "$dir/install.sh"
		echo "----------------------------------------------------------------------------------------------------"
		chmod +x "$dir/install.sh"
		echo "Running $dir/install.sh"

		# Run the install.sh script
		set -x
		"$dir/install.sh"
		set +x

		rm -rf "$dir/install.sh"
		echo "Test completed for $dir/"
	fi
done
