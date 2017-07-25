nbdropoutput
============

**nbdropoutput** is an executable that takes a stdin stream and filters out all of the output cells and execution_counts in a Jupyter Notebook (v4) so that it can be stored in a git repository

To do this follow the following steps:

1. Place the executable somewhere on your system.
    
        <Path-To-Executable>/nbdropoutput

2. Make it executable if needed
    
        chmod +x <Path-To-Executable>/nbdropoutput

3. Create an gitattributes file (.gitattributes) in the top level of repo with the following content

        *.ipynb    filter=dropoutput_ipynb

4. Configure the git repository to use the command on each commit

        git config filter.dropoutput_ipynb.clean <Path-To-Executable>/nbdropoutput

    NOTE: you need to do this after cloning a repository or make it a global configuration for git with the following command.

        git config --global filter.dropout_ipynb.clean <Path-To-Executable>/nbdropoutput

Adapted from References:

1. <https://stackoverflow.com/questions/18734739/using-ipython-notebooks-under-version-control>
2. <https://gist.github.com/pbugnion/ea2797393033b54674af>